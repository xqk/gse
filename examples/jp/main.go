package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/xqk/gse"
)

var (
	text = flag.String("text", `ひっそり远くから、もしかすると离（はな）し难（がた）いのか。
		黙々（もくもく）と静かに、もしかするととても価値（かち）があるのか。
		僕はまだここで待っている`, "単語セグメンテーションのテキスト")
)

func main() {
	flag.Parse()

	var seg gse.Segmenter
	// seg.LoadDict("../data/dict/dictionary.txt")
	err := seg.LoadDict("jp")
	fmt.Println("load dictionary error: ", err)

	fmt.Println(seg.Cut(*text, true))
	fmt.Println(seg.CutAll(*text))
	fmt.Println(seg.CutSearch(*text))

	fmt.Println(seg.String(*text, true))
	fmt.Println(seg.Slice(*text, true))

	segments := seg.Segment([]byte(*text))
	fmt.Println(gse.ToString(segments, true))

	text2 := []byte("运命は神の考えるものだ, 人间は人间らしく働ければそれ结构だ")
	segs := seg.Segment(text2)
	log.Println(gse.ToString(segs))
}
