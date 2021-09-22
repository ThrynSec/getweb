# getweb
A go package to easily handle downloading files and making API call

Basically :
* getweb.Download(a_great_file.txt, http://someurl.com/afile.txt) will download the file in the current folder and name it "a_great_file.txt". You can add a path to download it where you want.
* getweb.GetAPI(http://someurl.com/api/get_text) will return []bytes from the API.
