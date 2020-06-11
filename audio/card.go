/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package audio

import (
	"fmt"
	"sort"
	"strings"

	"pkg.deepin.io/lib/pulse"
	"pkg.deepin.io/lib/dbus1"
	bluez "github.com/linuxdeepin/go-dbus-factory/org.bluez"

)

type Card struct {
	Id            uint32
	Name          string
	ActiveProfile *Profile
	Profiles      ProfileList
	Ports         pulse.CardPortInfos
	core          *pulse.Card
}

type CardExport struct {
	Id    uint32
	Name  string
	Ports []CardPortExport
}

type CardPortExport struct {
	Name        string
	Description string
	Direction   int
}

func newCard(card *pulse.Card) *Card {
	var info = new(Card)
	info.core = card
	info.update(card)
	return info
}

func getCardName(card *pulse.Card) (name string) {
	propAlsaCardName := card.PropList["alsa.card_name"]
	propDeviceApi := card.PropList["device.api"]
	propDeviceDesc := card.PropList["device.description"]
	if propDeviceApi == "bluez" && propDeviceDesc != "" {
		name = propDeviceDesc
	} else if propAlsaCardName != "" {
		name = propAlsaCardName
	} else {
		name = card.Name
	}
	return
}

func (c *Card) update(card *pulse.Card) {
	//过滤掉蓝牙设备中被识别为computer的
	if strings.Contains(card.Name, "bluez") {
		cardType, err := getCardIcon(card.Name)
		if err != nil {
			return
		}
		if cardType == "computer" {
			return
		}
	}
	c.Id = card.Index
	c.Name = getCardName(card)
	c.ActiveProfile = newProfile(card.ActiveProfile)
	sort.Sort(card.Profiles)
	c.Profiles = newProfileList(card.Profiles)
	c.filterProfile(card)
	c.Ports = card.Ports
}

//从dbus中通过设备名字找到icon属性
func getCardIcon(cardName string) (cardType string, erro error) {
	systemBus, err := dbus.SystemBus()
	if err != nil {
		return "", err
	}
	n := strings.Split(cardName,".")
	var path string = "/org/bluez/hci0/dev_" + n[1]
	d, err := bluez.NewDevice(systemBus, dbus.ObjectPath(path))
	if err != nil {
		return "", err 
	}
	icon, err := d.Icon().Get(0)
	if err != nil {
		return "", err 
	}
	return icon, nil
}

func (c *Card) tryGetProfileByPort(portName string) (string, error) {
	profile, _ := c.Ports.TrySelectProfile(portName)
	if len(profile) == 0 {
		return "", fmt.Errorf("not found profile for port '%s'", portName)
	}
	return profile, nil
}

func (c *Card) filterProfile(card *pulse.Card) {
	var profiles ProfileList
	blacklist := profileBlacklist(card)
	for _, p := range c.Profiles {
		// skip unavailable and blacklisted profiles
		if p.Available == 0 || blacklist.Contains(p.Name) {
			// TODO : p.Available == 0 ?
			continue
		}
		profiles = append(profiles, p)
	}
	c.Profiles = profiles
}

type CardList []*Card

func newCardList(cards []*pulse.Card) CardList {
	var result CardList
	for _, v := range cards {
		result = append(result, newCard(v))
	}
	return result
}

func (cards CardList) string() string {
	var list []CardExport
	for _, cardInfo := range cards {
		var ports []CardPortExport
		for _, portInfo := range cardInfo.Ports {
			ports = append(ports, CardPortExport{
				Name:        portInfo.Name,
				Description: portInfo.Description,
				Direction:   portInfo.Direction,
			})
		}

		list = append(list, CardExport{
			Id:    cardInfo.Id,
			Name:  cardInfo.Name,
			Ports: ports,
		})
	}
	return toJSON(list)
}

func (cards CardList) get(id uint32) (*Card, error) {
	for _, info := range cards {
		if info.Id == id {
			return info, nil
		}
	}
	return nil, fmt.Errorf("invalid card id: %v", id)
}

func (cards CardList) add(info *Card) (CardList, bool) {
	card, _ := cards.get(info.Id)
	if card != nil {
		return cards, false
	}

	return append(cards, info), true
}

func (cards CardList) delete(id uint32) (CardList, bool) {
	var (
		ret     CardList
		deleted bool
	)
	for _, info := range cards {
		if info.Id == id {
			deleted = true
			continue
		}
		ret = append(ret, info)
	}
	return ret, deleted
}

func portAvailForCompare(v int) int {
	switch v {
	case pulse.AvailableTypeNo:
		return 0
	case pulse.AvailableTypeUnknow:
		return 1
	case pulse.AvailableTypeYes:
		return 2
	default:
		return -1
	}
}

func (cards CardList) getPassablePort(direction int) (cardId uint32,
	port *pulse.CardPortInfo) {

	for _, card := range cards {
		p := getPassablePort(card.Ports, direction)
		if p == nil {
			continue
		}

		if port == nil ||
			portAvailForCompare(p.Available) > portAvailForCompare(port.Available) ||
			(p.Available == port.Available && p.Priority > port.Priority) {
			cardId = card.Id
			port = p
		}
	}
	return
}

func getPassablePort(ports pulse.CardPortInfos, direction int) (port *pulse.CardPortInfo) {
	for idx := range ports {
		p := &ports[idx]
		if p.Direction != direction {
			continue
		}

		if port == nil ||
			portAvailForCompare(p.Available) > portAvailForCompare(port.Available) ||
			(p.Available == port.Available && p.Priority > port.Priority) {
			port = p
		}
	}
	return port
}
