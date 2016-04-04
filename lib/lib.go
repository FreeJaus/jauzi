package lib

import (
  "fmt"
  "os"
  "strings"
	"strconv"
  "errors"
	"regexp"
  "encoding/binary"
 	"encoding/gob"
  "path/filepath"
  "text/tabwriter"
  "github.com/fatih/color"
)

type Card struct {
  Name  string //16
  Copy  string //22
  Mods  string //16
  Date  []uint //3
  Data  []byte //1024
}

type Cards []*Card


type checkres struct {
 A,B string
 Match int
}


func EmptyCard() *Card { return &Card{ "", "", "", []uint{}, make([]byte,1024) } }

func NewCard(name,copy,mods string, date []uint, data []byte) *Card {
  return &Card{name,copy,mods,date,data}
}

func ID2Meta(n string) (name,copy,mods string, date []uint) {
  sn := strings.Split( filepath.Base(n)[0:len(filepath.Base(n))-len(filepath.Ext(n))] ,"_")

  if regexp.MustCompile(`[C]`).MatchString(sn[0]) {
   sn0 := strings.Split(sn[0],"C"); name = sn0[0]; copy = sn0[1];
  } else { name = sn[0] }

  date = make([]uint,3)

  if len(sn) > 1 {
   for i := 0; i < 3; i++ {
    x,_ := strconv.Atoi(sn[1][2*i:2*(i+1)])
    date[i]=uint(x)
  }}
  if len(sn) > 2 { mods = sn[2] }
  return
}

func (c *Card) Meta2ID() string {
 return Meta2ID(c.Name, c.Copy, c.Mods, c.Date)
}

func Meta2ID(name,copy,mods string, date []uint) (s string) {
  if len(copy)!=0 { s = name+"C"+copy+"_"; } else { s = name+"_" }
  for i := 0; i < 3; i++ { s = s+strconv.Itoa(int(date[i])) }
  if len(mods)!=0 { s = s+"_"+mods }
  return
}

func (c *Card) CmpMeta(b *Card) (r bool) {
 r = true
 if (c.Name!=b.Name) { r = false }
 if (c.Copy==b.Copy) { r = false }
 if (c.Mods==b.Mods) { r = false }
 for i := 0; i < 3; i++ { if c.Date[i]!=b.Date[i] { r = false } }
 return
}

func (c *Card) Read(p string, v bool) {
  if v { fmt.Println("  +", c.Meta2ID(), "'data' overwritten") }
  c.Data = *ReadMFD(p,1024)
}

func ReadCard(p string, v bool) *Card {
  n,c,m,d := ID2Meta(p)
  if v { fmt.Println("  R", n, c, m, d) }
  return NewCard( n, c, m, d, *ReadMFD(p,1024))
}

func ReadMFD(p string, size int) *[]byte{
  f, err := os.Open(p)
  if err != nil { fmt.Println(err); os.Exit(1) }
  defer f.Close()

  d := make([]byte,size)
  if err := binary.Read(f, binary.LittleEndian, &d); err != nil { fmt.Println("binary.Read failed:", err) }
  return &d
}

func (c *Card) Write(p string, v bool) { WriteCard(p,c,v) }

func WriteCard(p string, c *Card, v bool) {
  id := c.Meta2ID()

  stat,err := os.Stat(p)
  if err != nil { fmt.Println(err); os.Exit(1) }
	if stat.IsDir() { p = filepath.Join(p,id,"mfd") }
  if v { fmt.Println("  W", id) }

  WriteMFD(p,1024,&(c.Data))
}

func WriteMFD(p string, size int, d *[]byte) {
  f, err := os.Open(p)
  if err != nil { fmt.Println(err); os.Exit(1) }
  defer f.Close()

  if err := binary.Write(f, binary.LittleEndian, &d); err != nil { fmt.Println("binary.Write failed:", err) }
}

func (cs *Cards) Append(p string, v bool) {
  c := ReadCard(p,v)
  s := ""
  for i := 0; i < len(*cs); i++ {
   if c.CmpMeta((*cs)[i]) {
    for !regexp.MustCompile(`[sSoOnN]`).MatchString(s) {
     fmt.Printf("Card '%s' already present in DB. (S)kip/(O)verwrite/(N)ew?\n> ",c.Meta2ID())
      _,_ = fmt.Scanf("%s", &s)
      fmt.Println(s)
    }
    switch {
     case regexp.MustCompile(`[oO]`).MatchString(s): (*cs)[i] = c
     case regexp.MustCompile(`[nN]`).MatchString(s):
      m := true
      for m {
       c.Name = c.Name+"_N"
       m = false
       for j := 0; j < len(*cs); j++ { if c.Name==(*cs)[j].Name { m = true; } }
      }
    }
   }
  }
  if !regexp.MustCompile(`[sSoO]`).MatchString(s) { *cs = append(*cs, c ) }
}

func (cs *Cards) AppendWalk(dp string, rec,v bool) {
  fmt.Printf("- Appending cards in '%s' to the DB...\n",dp)
  if err := filepath.Walk(dp, func(p string, n os.FileInfo, err error) error {
    stat, err := os.Stat(p)
    if err != nil { return err }

    if stat.IsDir() && p != dp && !rec {
      fmt.Printf("  skipping dir '%s'\n", p[len(dp):len(p)])
      return filepath.SkipDir
    }

    if !stat.IsDir() && filepath.Ext(p)==".mfd"{ cs.Append(p,v) }
    return nil
  }); err != nil { fmt.Println("filepath.Walk failed:", err) }
}

func (c *Card) HexDump(b *[]bool) {
  for x := 0; x < 16; x++ {
    for i := 0; i < 4; i++ {
      d := (x*64)+(i*16)
      fmt.Printf("%.3X: ",d)
      lineDump(d,&c.Data,b)
      fmt.Println("")
    }
    fmt.Println("")
  }
}

func (c *Card) CmpHexDump(b *Card, l *[]bool) {
  for x := 0; x < 16; x++ {
    for i := 0; i < 4; i++ {
      d := (x*64)+(i*16)
      fmt.Printf("%.3X: ",d)
      lineDump(d,&c.Data,l)
      fmt.Print("|| ")
      lineDump(d,&b.Data,l)
      fmt.Println("")
    }
    fmt.Println("")
  }
}

func lineDump (d int, b *[]byte, c *[]bool) {
  for j := 0; j < 8; j++ { byteDump((*b)[d+j],(*c)[d+j]) }
  fmt.Print("| ")
  for j := 8; j < 16; j++ { byteDump((*b)[d+j],(*c)[d+j]) }
}

func byteDump (b byte, c bool) {
  red := color.New(color.FgRed).SprintFunc()
  green := color.New(color.FgGreen).SprintFunc()
  if c { fmt.Printf("%s",red(fmt.Sprintf("%.2x ",b)))
  } else { fmt.Printf("%s",green(fmt.Sprintf("%.2x ",b))) }
}

func NewCardDB() Cards { return Cards{} }

func (cs *Cards) Extract(ids []int) Cards {
  fmt.Printf("- Extracting %d cards to a new DB:\n", len(ids))
  ecs := NewCardDB()
  for i := 0; i < len(ids); i++ {
   ecs = append(ecs,(*cs)[ids[i]])
   fmt.Println("  +",(*cs)[ids[i]].Name)
  }
  return ecs
}

func ReadDB(p string, v bool) Cards {
  fmt.Printf("- Reading DB from '%s'\n",p)
  cs := NewCardDB()
  f, err := os.Open(p)
  if err != nil { fmt.Println(err); os.Exit(1); }
  defer f.Close()
  dec := gob.NewDecoder(f)
  if err := dec.Decode(&cs); err != nil { fmt.Println(err); os.Exit(1); }
  if v {
   fmt.Println("-", len(cs),"cards read:")
   for i := 0; i < len(cs); i++ { fmt.Printf("  %*d. '%s'\n",3,i,(*cs[i]).Name) }
  }
  return cs
}

func (cs *Cards) WriteDB(p string) {
  fmt.Printf("- Saving DB to '%s'\n",p)
  f, err := os.Create(p)
  if err != nil { fmt.Println(err); os.Exit(1); }
  defer f.Close()
  enc := gob.NewEncoder(f)
 	enc.Encode(cs)
}

func (cs *Cards) ParseList(ls *string) []int {
  fmt.Println("- Parsing list...")

  el := strings.Split(*ls,",")
  ids := []int{}

  for i := 0; i < len(el); i++ {
   if regexp.MustCompile(`^(0|[1-9][0-9]*)$`).MatchString(el[i]) {
     k,err := strconv.Atoi(el[i])
     if err != nil { err = errors.New(fmt.Sprintf("strconv.Atoi failed: %s", err))
     } else if k>(len(*cs)-1) {
      err = errors.New(fmt.Sprintf("index %d out of bounds [0,%d]", k, len(*cs)-1))
     }
     if err != nil { fmt.Printf("'%s' skipped: %s\n",el[i],err)
     } else {
      ids = append(ids,k)
     }
   } else {
    m := true
    for j := 0; j < len(*cs); j++ {
     if (*cs)[j].Name==el[i] { ids = append(ids,j); m = false; break }
    }
    if m { fmt.Printf("  '%s' skipped: name not found in DB\n",el[i]) }
   }
  }

  fmt.Printf("- %d valid ids parsed out of %d\n",len(ids),len(el))

  return ids
}

func (cs *Cards) CheckEq(pr bool) (*[]checkres, []int) {
  cmps := make([]checkres,(10*9)/2)
  var c,k int = 0,0
  min := []int{1024,0}
  for i := 0; i < len(*cs); i++ {
   for j := i+1; j < len(*cs); j++ {
    c = 0
    for x := 0; x < 1024; x++ { if (*cs)[i].Data[x]==(*cs)[j].Data[x] { c++; }  }
    cmps[k].A=(*cs)[i].Name
    cmps[k].B=(*cs)[j].Name
    cmps[k].Match=c
    if min[0]>c {min=[]int{c,k}}
    k++
   }
  }
  if pr {
    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 0, 8, 1, '\t', 0)
    k=0
    for i := 0; i < len(*cs); i++ {
     for j := i+1; j < len(*cs); j++ {
      fmt.Fprintln(w,fmt.Sprintf(" %.2d: %s\t%s\t\t%d\t%.2f\t",k,cmps[k].A,cmps[k].B,cmps[k].Match,float32(cmps[k].Match)*100/1024))
      k++
     }
    }
    w.Flush()
  }
  return &cmps,min
}

func (cs *Cards) NotCommon (pr bool) (*[]bool,*[]byte) {
  cmps := make([]bool,1024)
  for i := 0; i < len(*cs); i++ {
   for j := i+1; j < len(*cs); j++ {
    for x := 0; x < 1024; x++ { if (*cs)[i].Data[x]!=(*cs)[j].Data[x] { cmps[x]=true; }  }
   }
  }
  if pr {
   c :=  NewCard("common","","",[]uint{},make([]byte,1024))
   for x := 0; x < 1024; x++ { if !cmps[x] { c.Data[x]= (*cs)[0].Data[x]} }
   if len(*cs)==2 { (*cs)[0].CmpHexDump((*cs)[1], &cmps) } else { c.HexDump(&cmps) }
   return &cmps,&c.Data;
  }  else {
   return &cmps,&[]byte{}
  }
}
