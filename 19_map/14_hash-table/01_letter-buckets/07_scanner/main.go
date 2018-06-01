package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main(){
	input := `Raskolnikov was violently agitated. Of course, it was all ordinary
	youthful talk and thought, such as he had often heard before in
	different forms and on different themes. But why had he happened to hear
	such a discussion and such ideas at the very moment when his own brain
	was just conceiving... _the very same ideas_? And why, just at the
	moment when he had brought away the embryo of his idea from the old
	woman had he dropped at once upon a conversation about her? This
	coincidence always seemed strange to him. This trivial talk in a tavern
	had an immense influence on him in his later action; as though there had
	really been in it something preordained, some guiding hint....`

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	// the Scan method returns a bool for whether or not a word exists or not
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "reading-input:", err)
	}
}