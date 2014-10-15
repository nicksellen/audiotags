/***************************************************************************
   copyright            : (C) 2014 by Nick Sellen
   email                : code@nicksellen.co.uk
***************************************************************************/

/***************************************************************************
 *   This library is free software; you can redistribute it and/or modify  *
 *   it  under the terms of the GNU Lesser General Public License version  *
 *   2.1 as published by the Free Software Foundation.                     *
 *                                                                         *
 *   This library is distributed in the hope that it will be useful, but   *
 *   WITHOUT ANY WARRANTY; without even the implied warranty of            *
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU     *
 *   Lesser General Public License for more details.                       *
 *                                                                         *
 *   You should have received a copy of the GNU Lesser General Public      *
 *   License along with this library; if not, write to the Free Software   *
 *   Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  *
 *   USA                                                                   *
 ***************************************************************************/

package main

import (
	"fmt"
	"github.com/nicksellen/audiotags"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("pass path to file")
		return
	}

	props, audioProps, err := audiotags.Read(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range props {
		fmt.Printf("%s %s\n", k, v)
	}

	fmt.Printf("length %d\nbitrate %d\nsamplerate %d\nchannels %d\n",
		audioProps.Length, audioProps.Bitrate, audioProps.Samplerate, audioProps.Channels)

}
