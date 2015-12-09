package zhtool

import (
  "testing"
)

func Test_Chs2Cht(t *testing.T) {
  s := "遇到一个诡异Bug，每逢周三就崩溃"
  if s != Cht2Chs(Chs2Cht(s)) {
    t.Error("Error")
  }
}
