//          ATCache  Copyright (C) 2018  AnimeTwist
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package redis

import (
	"github.com/go-redis/redis"
	"log"
	"net"
)

var Client *redis.Client

func Load(host, port, password string, DB int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(host, port),
		Password: password,
		DB:       DB,
	})

	if r, err := Client.Ping().Result(); err != nil {
		return err
	} else {
		log.Println("Redis Ping:", r)
	}

	return nil
}
