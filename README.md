# getweb
A go package to easily handle downloading files and making API call

Basically :
* getweb.Download(http://someurl.com/afile.txt, a_great_file.txt) will download the file in the current folder and name it "a_great_file.txt"
* getweb.GetAPI(http://someurl.com/api/get_text) will return []bytes from the API
