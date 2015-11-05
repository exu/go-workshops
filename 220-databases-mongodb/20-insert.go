package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"time"
)

type Person struct {
	Name        string
	Phone       int64
	Description string
}

const DESCRIPTION = `Lorem ipsum dolor sit amet, consectetur adipisicing elit. Proin nibh augue, suscipit a, scelerisque sed, lacinia in, mi. Cras vel lorem. Etiam pellentesque aliquet tellus. Phasellus pharetra nulla ac diam. Quisque semper justo at risus. Donec venenatis, turpis vel hendrerit interdum, dui ligula ultricies purus, sed posuere libero dui id orci. Nam congue, pede vitae dapibus aliquet, elit magna vulputate arcu, vel tempus metus leo non est. Etiam sit amet lectus quis est congue mollis. Phasellus congue lacus eget neque. Phasellus ornare, ante vitae consectetuer consequat, purus sapien ultricies dolor, et mollis pede metus eget nisi. Praesent sodales velit quis augue. Cras suscipit, urna at aliquam rhoncus, urna quam viverra nisi, in interdum massa nibh nec erat.
Mauris id urna. Donec adipiscing. Cras luctus est id urna. Sed vitae enim. Donec sit amet magna molestie quam varius dictum. Morbi laoreet. Maecenas elit lacus, congue eu, rutrum ac, condimentum et, risus. Donec erat. In ornare placerat nisl. Aliquam eget metus non massa tristique luctus. Integer id justo. Fusce rutrum. Nunc volutpat magna pretium massa.
Proin gravida tellus vitae dui. Maecenas tempor. Nulla eros mauris, aliquet quis, posuere ut, ultrices non, quam. Pellentesque at odio semper justo placerat convallis. Pellentesque diam. Sed mattis, pede ut ultricies commodo, velit magna auctor leo, at ultricies velit risus id erat. Donec et eros in urna interdum vulputate. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nunc fermentum nibh id ante. Aliquam felis augue, gravida non, venenatis sed, congue et, orci. Aenean scelerisque tortor sed magna. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nullam hendrerit, mi et mattis lobortis, nibh ante vehicula magna, interdum porta massa mauris vitae odio. Integer aliquam arcu non enim. Fusce felis metus, dignissim in, molestie vel, pellentesque a, odio.
Lorem ipsum dolor sit amet, consectetur adipisicing elit. Proin nibh augue, suscipit a, scelerisque sed, lacinia in, mi. Cras vel lorem. Etiam pellentesque aliquet tellus. Phasellus pharetra nulla ac diam. Quisque semper justo at risus. Donec venenatis, turpis vel hendrerit interdum, dui ligula ultricies purus, sed posuere libero dui id orci. Nam congue, pede vitae dapibus aliquet, elit magna vulputate arcu, vel tempus metus leo non est. Etiam sit amet lectus quis est congue mollis. Phasellus congue lacus eget neque. Phasellus ornare, ante vitae consectetuer consequat, purus sapien ultricies dolor, et mollis pede metus eget nisi. Praesent sodales velit quis augue. Cras suscipit, urna at aliquam rhoncus, urna quam viverra nisi, in interdum massa nibh nec erat.
Mauris id urna. Donec adipiscing. Cras luctus est id urna. Sed vitae enim. Donec sit amet magna molestie quam varius dictum. Morbi laoreet. Maecenas elit lacus, congue eu, rutrum ac, condimentum et, risus. Donec erat. In ornare placerat nisl. Aliquam eget metus non massa tristique luctus. Integer id justo. Fusce rutrum. Nunc volutpat magna pretium massa.
Proin gravida tellus vitae dui. Maecenas tempor. Nulla eros mauris, aliquet quis, posuere ut, ultrices non, quam. Pellentesque at odio semper justo placerat convallis. Pellentesque diam. Sed mattis, pede ut ultricies commodo, velit magna auctor leo, at ultricies velit risus id erat. Donec et eros in urna interdum vulputate. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nunc fermentum nibh id ante. Aliquam felis augue, gravida non, venenatis sed, congue et, orci. Aenean scelerisque tortor sed magna. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nullam hendrerit, mi et mattis lobortis, nibh ante vehicula magna, interdum porta massa mauris vitae odio. Integer aliquam arcu non enim. Fusce felis metus, dignissim in, molestie vel, pellentesque a, odio.
Lorem ipsum dolor sit amet, consectetur adipisicing elit. Proin nibh augue, suscipit a, scelerisque sed, lacinia in, mi. Cras vel lorem. Etiam pellentesque aliquet tellus. Phasellus pharetra nulla ac diam. Quisque semper justo at risus. Donec venenatis, turpis vel hendrerit interdum, dui ligula ultricies purus, sed posuere libero dui id orci. Nam congue, pede vitae dapibus aliquet, elit magna vulputate arcu, vel tempus metus leo non est. Etiam sit amet lectus quis est congue mollis. Phasellus congue lacus eget neque. Phasellus ornare, ante vitae consectetuer consequat, purus sapien ultricies dolor, et mollis pede metus eget nisi. Praesent sodales velit quis augue. Cras suscipit, urna at aliquam rhoncus, urna quam viverra nisi, in interdum massa nibh nec erat.
Mauris id urna. Donec adipiscing. Cras luctus est id urna. Sed vitae enim. Donec sit amet magna molestie quam varius dictum. Morbi laoreet. Maecenas elit lacus, congue eu, rutrum ac, condimentum et, risus. Donec erat. In ornare placerat nisl. Aliquam eget metus non massa tristique luctus. Integer id justo. Fusce rutrum. Nunc volutpat magna pretium massa.
Proin gravida tellus vitae dui. Maecenas tempor. Nulla eros mauris, aliquet quis, posuere ut, ultrices non, quam. Pellentesque at odio semper justo placerat convallis. Pellentesque diam. Sed mattis, pede ut ultricies commodo, velit magna auctor leo, at ultricies velit risus id erat. Donec et eros in urna interdum vulputate. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nunc fermentum nibh id ante. Aliquam felis augue, gravida non, venenatis sed, congue et, orci. Aenean scelerisque tortor sed magna. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Nullam hendrerit, mi et mattis lobortis, nibh ante vehicula magna, interdum porta massa mauris vitae odio. Integer aliquam arcu non enim. Fusce felis metus, dignissim in, molestie vel, pellentesque a, odio.
`

func main() {
	session, err := mgo.Dial("localhost:7702")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("goworkshops_insert").C("people")
	c.DropCollection()

	start := time.Now()
	recordsCount := 100000

	for i := 0; i < recordsCount; i++ {
		err = c.Insert(&Person{fmt.Sprintf("Ale %d", i), rand.Int63(), DESCRIPTION})

		if err != nil {
			log.Fatal(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Added %d into goworkshops_insert.people collection in %s\n", recordsCount, elapsed)
}
