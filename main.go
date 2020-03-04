package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"crypto/sha256"
	"encoding/hex"
	"encoding/binary"
	"bytes"
	//"strings"
)

type Timestamp struct {
	Operator string // [0]
	Prefix string   // [1]
	Postfix string  // [2]
}

// This function should walk through the timestamps and verify message against merkleRoot
// Hints: use crypto/sha256 and encoding/hex. message is big-endian while merkleRoot is little-endian.
func VerifyHash(timestamps []Timestamp, message string, merkleRoot string) bool {

	// big-endian to little-endian convert

	treeHashes := []string{}
	for  k := 0; k < len(timestamps); k++ {
		//b := []byte(timestamps[k].Prefix+message+timestamps[k].Postfix)
		//newHash := sha256.Sum256(b)
		//treeHashes = append(treeHashes, hex.EncodeToString(newHash[:]))
		
	    	//fmt.Printf("%s \n", timestamps[k].Prefix+msg_l_endian+timestamps[k].Postfix)
	    	//fmt.Printf("\n")
		treeHashes = append(treeHashes, timestamps[k].Prefix+message+timestamps[k].Postfix)
		
	}

	// If tree length become odd
	if len(treeHashes)%2 != 0 {
		treeHashes = append(treeHashes, treeHashes[len(treeHashes)-1])
	}
	tempHash := []string{}
	for {
	    
	    tempHash = nil
	    fmt.Printf("%s \n", treeHashes)
	    i:= 1
	    fmt.Printf("\n")
	    l := len(treeHashes)
            for  i <= l {
		if i % 2 == 0 {
			//var bin_buf bytes.Buffer
			//binary.Write(&bin_buf, binary.BigEndian, treeHashes[i-2]+treeHashes[i-1])

			b := []byte(treeHashes[i-2]+treeHashes[i-1])
		    	newHash := sha256.Sum256(b)
			
			
		    	fmt.Printf("%s \n", treeHashes[i-2]+treeHashes[i-1])
		    	fmt.Printf("\n")
			tempHash = append(tempHash, hex.EncodeToString(newHash[:]))
		}
		i++
		
	    }
            if len(treeHashes) < 2 {
                break
            }
	    treeHashes = nil
	    treeHashes = tempHash
        }
	
	var root_b_binary bytes.Buffer
	binary.Write(&root_b_binary, binary.BigEndian, merkleRoot)
	root_b_hash := sha256.Sum256(root_b_binary.Bytes())
	root_b_endian := hex.EncodeToString(root_b_hash[:])
	if treeHashes[0] == root_b_endian {
		return true
	} else {
		return false
	}
 
}


func main(){
	msg := "b4759e820cb549c53c755e5905c744f73605f8f6437ae7884252a5f204c8c6e6"
	merkleRoot := "f832e7458a6140ef22c6bc1743f09610281f66a1b202e7b4d278b83de55ef58c"

	filePath := "./bag/timestamp.json"
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	tuples := [][]string{}
	
	timestamps := []Timestamp{}

	// Convert byte slice to slice of tuples
	json.Unmarshal(dat, &tuples)

   	for  i := 0; i < len(tuples); i++ {
	   var temp = tuples[i];
		var opt, prx, ptx string
      	   for j := 0; j < len(temp); j++ {
		if (j == 0) {
		   opt = temp[j];
		} else if (j == 1) {
		   prx = temp[j];
		} else if (j == 2){
		   ptx = temp[j];
		}
      	  }
	  
	  t := Timestamp{Operator: opt, Prefix: prx, Postfix: ptx};
	  timestamps = append(timestamps, t)
   	}

	if VerifyHash(timestamps, msg, merkleRoot) == true {
		fmt.Println("CORRECT!")
	} else {
		fmt.Println("INCORRECT!")
	}
}
