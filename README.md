# go-su
go-su는 go slice를 쉽게 사용하기 위한 utility 패키지입니다.
# 시작하기
아래와 같이 패키지를 설치 합니다.
```
$ go get github.com/swkwon/go-su@latest
```
## 객체 생성
su를 사용하기 위해서 먼저 New함수로 인스턴스를 생성합니다.
```
    mySlice := su.New[int]()
```
이미 생성된 slice로도 생성이 가능합니다.
```
    myIntSlice := []int{1,2,3}
    mySlice := su.New(myIntSlice)
```
여러개의 slice를 하나로 만들면서 생성할 수 있습니다.
```
    mySlice := su.New([]int{1,2,3}, []int{4,5,6})
```
## 삽입
삽입에는 두가지 방법이 있습니다.
### Append
Append는 데이터를 뒤에 삽입합니다.
```
    mySlice := su.New[int]()
    mySlice.Append(1, 2, 3)
    // 또는
    mySlice.Append([]int{7, 8, 9})
    // output
    // [1, 2, 3, 7, 8, 9]
```
### Prepend
Prepend는 데이터를 앞에 삽입합니다.
```
    mySlice := su.New[int]()
    mySlice.Prepend(1, 2, 3)
    // 또는
    mySlice.Prepend([]int{7, 8, 9})
    // output
    // [7, 8, 9, 1, 2, 3]
```
## 제거
제거에는 4가지 방법이 있습니다.
### Remove
Remove는 원하는 slice의 인덱스를 입력하면 해당 요소가 제거 됩니다.
```
    mySlice := su.New([]int{10, 20, 30})
    mySlice.Remove(0)
    // output
    // [20, 30]
```
### RemoveRange
RemoveRange는 연속된 범위를 제거합니다. 파라메터로 start, end 인덱스를 입력하는데 인덱스를 포함하여 제거 합니다.
```
    mySlice := su.New([]int{10, 20, 30, 40, 50})
    mySlice.RemoveRange(1, 2)
    // output
    // [10, 40, 50]
```
### RemoveAll
RemoveAll은 모두 삭제 합니다.
```
    mySlice := su.New([]int{10, 20, 30, 40, 50})
    mySlice.RemoveAll()
    // output
    // []
```
### 이터레이터를 이용한 Remove
이터레이터를 이용하여 loop를 돌며 요소를 삭제할 수도 있습니다.
```
    mySlice := su.New([]int{1, 2, 3, 4, 5, 6})
    itr := mySlice.Iterator()
    for itr.MoveNext() {
        if _, v := itr.Current(); v % 2 == 0 {
            itr.Remove()
        }
    }
    // output
    // [1, 3, 5]
```
## 데이터 검색, 가져오기, 저장하기
slice내에 있는 값을 찾고, 값을 가져오고, 다른 값으로 덮어씌울 수 있습니다.
```
	mySlice := su.New([]int{1, 2, 3, 4})

    // index는 3 입니다.
	index := mySlice.IndexOf(func(value int) bool {
		return value == 4
	})
    
    // value는 4 입니다.
    value := mySlice.Get(3)
    
    // 4번째 데이터의 값은 10이 됩니다.
    mySlice.Set(3, 10)

    // output
    // [1, 2, 3, 10]
```