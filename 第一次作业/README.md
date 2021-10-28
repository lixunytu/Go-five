## local tree
.
├── app
│   ├── dao
│   │   └── user_dao.go
│   ├── databases
│   │   └── mysql.go
│   └── service
├── go.mod
├── go.sum
└── main.go

## 题目
  1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 答
  要告知上层，因为上层使用的时候 nil 和错误是一样的 ，无法使用，抛给上层统一处理
  如果需要知道是什么错误，或者需要具体的判断可以使用errors.Is 去判断什么错误
