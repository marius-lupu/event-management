/*
	Copyright (c) 2022 Marius Lupu

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"fmt"
	"net/http"
)

// verifyCredentials will check if the Authorization header has the right username and password
func verifyCredentials(r *http.Request) error {
	givenUser, givenPass, ok := r.BasicAuth()
	if !ok {
		return fmt.Errorf("no basic authorization was provided, please provide one")
	}
	if givenUser != appUsername || givenPass != appPassword {
		return fmt.Errorf("username and/or password are invalid")
	}
	return nil
}
