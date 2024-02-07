Fork of [Switch Library Manager](https://github.com/giwty/switch-library-manager) created by giwty

##### CHANGES
- TITLES_JSON_URL and VERSIONS_JSON_URL now are configurable in settings.json
- TitleID in library tab is upper case now
- Corrected button from "export to scv" to "export to CSV"
- "Updates" tab now is "Missing updates"
- "DLC" tab now is "Missing DLC"

##### FIXES
- New keys are now well calculated
- First 13 ID hexadecimal chars are used to get title information insted of 12. It caused "Duplicate game base" on games with the same first ID 12 hexadecimal chars
- File names starting with a dot ('.') now are accepted
- Last title row is now exported

##### IMPROVEMENTS
- Added "ignore_demos" option in 'settings.json' . All titles with field "isDemo" set to true in "titles.json" are ignored
- Added "blacklist.json" support. If present, all titles/dlc in this file are ignored. Ideal to avoid those games you'll never want in your list. Fields -> "Id" (mandatory), "Name", "Region", "Reason"
- Added "whitelist.json" support. If present, all content included in this file is always listed. Ideal to add cart only games, homebrew or to redefine an already existing title listed in 'titles.json'. File struct the same 'titles.json' but only ID is mandator
- Added TITLEID, REGION and TYPE filter to "LIBRARY" tab
- Added TITLEID and REGION filter to "MISSING GAMES" tab
- Added TITLEID filter to "MISSING DLC" tab
- Added direct access to file update if exists. Just press CONTROL key on keyboard and then click on UPDATE or VERSION value
- New export method in MISSING DLC tab. Now creates a row for every missing dlc. The old method is also kept but has been renamed to "export to CSV (compact)"

##### KNOWN ISSUES
- Download 'titles.json' and 'versions.json' from tinfoil not working. Get an x509 certificate error. Tried to solve it accepting all certificates, but I get an 404 - NOT FOUND error. I recommed to download them manually or change TITLES_JSON_URL and VERSIONS_JSON_URL


## Usage
##### Windows
- Extract the zip file
- Double click the Exe file
- If you want to use command line mode, update the settings.json with `'GUI':false`
    - Open `cmd`
    - Run `switch-library-manager.exe`
    - Optionally -f `X:\folder\containing\nsp\files"`
    - Optionally add  `-r` to recursively scan for nested folders
    - Edit the settings.json file for additional options

 
##### macOS or Linux
- Extract the zip file
- Double click the App file
- If you want to use command line mode, update the settings.json with `'GUI':false`
    - Open your Terminal
    - `cd` to the folder containing `switch-library-manager`
    - `chmod +x switch-library-manager` to make it executable
    - Run `./switch-library-manager'
    - Optionally -f `X:\folder\containing\nsp\files"`
    - Optionally add  `-r` to recursively scan for nested folders
    - Edit the settings.json file for additional options

## Building
- Install and setup Go
- Clone the repo: `git clone https://github.com/pedro-dro/switch-library-manager.git`
- Get the bundler `go get -u github.com/asticode/go-astilectron-bundler/...`
- Install bundler `go install github.com/asticode/go-astilectron-bundler/astilectron-bundler`
- Copy bundler binary to the source folder `cd switch-library-manager` and then `mv $HOME/go/bin/astilectron-bundler .`
- Execute `./astilectron-bundler`
- Binaries will be available under output

#### Thanks
- To @giwty for his great job
- This program relies on [blawar's titledb](https://github.com/blawar/titledb), to get the latest titles and versions.
