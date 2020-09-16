package kata

func Race(v1, v2, g int) [3]int {
  if v1 >= v2 {
      return [3]int{-1, -1, -1}
    }
  //Consider: d_1 = g + v_1 * t, d_2 = v_2 * t
  // B catches A when d_1 = d_2,
  //                  g + v_1 * t = v_2 * t
  //                            g = v_2 * t - v_1 * t
  //                            t = g /(v_2-v_1)
  const secondsPerHour = 3600
  // Convert from feet per hour to feet per second
  
  seconds := g*secondsPerHour/(v2-v1)
  hours := seconds/secondsPerHour
  seconds -= hours * secondsPerHour
  minutes := seconds / 60
  seconds -= minutes * 60
  
  return [3]int{hours, minutes, seconds}
}