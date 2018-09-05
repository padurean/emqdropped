package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Message ...
type Message struct {
	// Client  string
	Topic   string
	Payload string
}

// TopicMessages ...
type TopicMessages struct {
	Topic string
	Msgs  []Message
	Count int
}

// MessagesPerTopic ...
type MessagesPerTopic map[string]TopicMessages

func extractMsg(l string) *Message {
	// client := strings.Split(strings.Split(l, "Client(")[1], "):")[0]
	// client = strings.Split(client, ":")[0]
	topic := strings.Split(strings.Split(l, "Topic=")[1], ", ")[0]
	payload := strings.Split(strings.Split(l, `Payload=<<"`)[1], `">>`)[0]
	return &Message{
		// Client:  client,
		Topic:   topic,
		Payload: payload,
	}
}

func addMsgToMap(m map[Message]int, l string) {
	msg := extractMsg(l)
	i, ok := m[*msg]
	if ok {
		m[*msg] = i + 1
	} else {
		m[*msg] = 1
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print(`Specify emq log path - e.g. "emqdropped ./emq.sample.log"`)
		return
	}
	var showPayload = false
	if len(os.Args) == 3 {
		sp, err := strconv.ParseBool(os.Args[2])
		if err != nil {
			fmt.Println("ERROR: 2nd argument", os.Args[2], "not understood. Expected true/false (to show/hide dropped messages payload).")
		}
		showPayload = sp
	}
	path := os.Args[1]
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var receives = make(map[Message]int)
	var sends = make(map[Message]int)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "RECV PUBLISH") {
			addMsgToMap(receives, line)
		} else if strings.Contains(line, "SEND PUBLISH") {
			addMsgToMap(sends, line)
		}
	}

	check(scanner.Err())

	var diff = make(map[Message]int)
	ctr := 0
	for r, c := range receives {
		ctr += c
		sc, ok := sends[r]
		if !ok {
			diff[r] = c
		} else if sc < c {
			diff[r] = c - sc
		}
	}
	fmt.Println(ctr, "published")

	cts := 0
	for _, c := range sends {
		cts += c
	}
	fmt.Println(cts, "delivered")

	ctd := 0
	var diffPerTopic = make(MessagesPerTopic)
	for d, c := range diff {
		ctd += c
		tms, ok := diffPerTopic[d.Topic]
		if !ok {
			diffPerTopic[d.Topic] = TopicMessages{Topic: d.Topic, Msgs: []Message{d}, Count: c}
		} else {
			tms.Msgs = append(tms.Msgs, d)
			tms.Count += c
			diffPerTopic[d.Topic] = tms
		}
	}
	fmt.Println(ctd, "dropped")

	for t, tms := range diffPerTopic {
		fmt.Println("\t", t, ":", tms.Count, "dropped")
		if showPayload {
			for _, msg := range tms.Msgs {
				fmt.Println("\t\t", msg.Payload)
			}
		}
	}
	fmt.Println("FINISHED.")
}
