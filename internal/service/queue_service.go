/*
 *  * Copyright (c) 2025 @umfaka
 *  *
 *  * This program is free software: you can redistribute it and/or modify
 *  * it under the terms of the GNU Affero General Public License as
 *  * published by the Free Software Foundation, either version 3 of the
 *  * License, or (at your option) any later version.
 *  *
 *  * 本程序为自由软件：你可以在遵守 GNU Affero 通用公共许可证（AGPL）第3版或（由你选择）任何更高版本的前提下重新发布和/或修改本程序。
 *  *
 *  * This program is distributed in the hope that it will be useful,
 *  * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  * GNU Affero General Public License for more details.
 *  *
 *  * 本程序以“按现状”提供，不提供任何明示或暗示的保证，包括但不限于适销性或特定用途适用性的保证。
 *  * 有关详细信息，请参阅 GNU Affero 通用公共许可证。
 *  *
 *  * You should have received a copy of the GNU Affero General Public License
 *  * along with this program. If not, see <https://www.gnu.org/licenses/>.
 *  *
 *  * 您应该已经收到一份 GNU Affero 通用公共许可证的副本；如果没有，请访问：<https://www.gnu.org/licenses/>
 *  *
 *  * Team:   @umfaka
 *  * Author: @uncledawen
 *  * Project: annona_client
 *  * File: queue_service.go
 *  * LastModified: 2025/5/26 23:40
 */

package service

import (
	"fmt"
	"github.com/AnnonaOrg/annona_client/internal/db_data"
)

var noUsernameChatIDQueue = &db_data.FIFOMap{}

func init() {
	noUsernameChatIDQueue.Init()
}
func CheckNoUsernameChatIDQueue(chatID int64) bool {
	key := fmt.Sprintf("%d", chatID)
	if _, ok := noUsernameChatIDQueue.Get(key); ok {
		return true
	} else {
		return false
	}
}

func SetNoUsernameChatIDQueue(chatID int64) {
	key := fmt.Sprintf("%d", chatID)
	noUsernameChatIDQueue.Set(key, true)
}
func RemoveOldestNoUsernameChatIDQueue() {
	if count := noUsernameChatIDQueue.Count(); count > 50 {
		noUsernameChatIDQueue.RemoveOldest()
	}
}
