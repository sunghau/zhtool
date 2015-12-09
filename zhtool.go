package zhtool

//Chs2Cht translate simplified chinese word to traditional chinese word
func Chs2Cht(str string) string {
  res := ""
  for _, s := range str {
    c := string(s)

    if len(S2T[c]) != 0 {
      res += S2T[c]
    } else if len(S2T2[c]) != 0 {
      res += S2T2[c]
    } else {
      res += c
    }
  }

  return res
}

//Cht2Chs translate traditional chinese word to simplified chinese word
func Cht2Chs(str string) string {
  res := ""
  for _, s := range str {
    c := string(s)
    if len(T2S[c]) != 0 {
      res += T2S[c]
    } else if len(T2S2[c]) != 0 {
      res += T2S2[c]
    } else {
      res += c
    }
  }

  return res
}
