1. to run tests , the test file must have the same name of the main file with sufix "_test"

2. to run tests, go to the folder where the main and the text files reside and type "go test"

3. to a full view of covering of main file tests, type "go test -cover"

4. to a better solution than 3, type "go test -coverprofile=coverage.out && go tool cover -html=coverage.out"