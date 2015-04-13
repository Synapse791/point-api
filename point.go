package main

import (

)

type Point struct {
    Title       string  `json:"title"`
    Date        int32   `json:"date"`
    User        string  `json:"user"`
    Mega        bool    `json:"mega"`
}

type Points []Point
