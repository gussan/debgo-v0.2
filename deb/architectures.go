/*
   Copyright 2013 Am Laher

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

package deb

//A processor architecture (ARM/x86/AMD64) - as named by Debian. At this stage: i386, armel, amd64
type Architecture string

 const (
	Arch_i386 Architecture = "i386"
	Arch_armel Architecture = "armel" //TODO armhf
	Arch_amd64 Architecture = "amd64"
)