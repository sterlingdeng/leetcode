function quickSort(array) {
  // Write your code here.
  quickSortHelper(array, 0, array.length - 1);
  return array;
}

function quickSortHelper(array, low, high) {
  if (low < high) {
    const partitionIdx = partition(array, low, high);
    // do left side
    quickSortHelper(array, low, partitionIdx - 1);
    quickSortHelper(array, partitionIdx + 1, high);
  }
}

function partition(array, low, high) {
  const partitionValue = array[high];
  let pIdx = low;
  // ensure values that are below the partitionValue is left of the partition Idx
  for (let i = low; i < high; i++) {
    const num = array[i];
    if (num < partitionValue) {
      swap(array, i, pIdx);
      pIdx++;
    }
  }
  swap(arr, pIdx, high);
  return pIdx;
}

function swap(arr, a, b) {
  const temp = arr[a];
  arr[a] = arr[b];
  arr[b] = temp;
}

const arr = [1, 2];
quickSort(arr);
console.log(arr);
