function meetingRooms(arr) {
   arr.sort((a, b) => a[0] - b[0]);
   for (let i = 0; i < arr.length - 1; i++) {
      if (arr[i][1] > arr[i+1][0]) {
         return false
      }
   }
   return true
}

