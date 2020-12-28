
package main


// after importing package whose path has multiple components
// the pckage is refered using the name that comes from the last component
import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
)


// composite literals
// instatiate composite types from sequence of element values
// instantiate slice
var palette = []color.Color{color.White, color.Black}

// const declarations
// these values are fixed at compile time
// must be string, number or boolean
const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main(){
    liss_handler := func(w http.ResponseWriter, r *http.Request){
        lissajous(w)
    }
    http.HandleFunc("/", liss_handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
    const (
        cycles = 5    // number of complete x oscillator cycles
        res = 0.001   // angular resolution
        size = 100    // image canvas convers [-size..+size]
        nframes = 64  // number of animation frames
        delay = 8      // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative freq of y oscillator
    // instantiate struct
    // struct is a group of values called fields
    // create a struct of the type gif.GIF whose LoopCount is set to nframes
    // all aother fields values will have the 'zero' value for their type
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase diff
    // outer loops runs for the number of frames
    // it creates a 201x201 image with a palette of two colours
    for i := 0; i < nframes; i++{
        rect := image.Rect(0,0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        // set some pixels of the image to black
        // this loops runs for 5 oscillations
        // x & y varies from -100 to 100 in a sinusoid fashion
        for t := 0.0; t < cycles*2*math.Pi; t+= res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            // set pixel values
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                blackIndex)
        }
        phase += 0.1

        // indiviudual elements of the struct anim can be accessed by 
        // using the dot notation
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
