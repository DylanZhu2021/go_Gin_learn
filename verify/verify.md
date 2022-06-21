**学习Gin框架的参数验证先了解一下[validator库](https://github.com/go-playground/validator)**

**gin框架是使用validator.v10这个库来进行参数验证的，我们先来看看这个库的使用。**

[资料来源](https://segmentfault.com/a/1190000023725115)

### 然后先写一个简单的示例：

```go
package main

import (
    "fmt"

    "github.com/go-playground/validator/v10"
)

type User struct {
    Username string `validate:"min=6,max=10"`
    Age      uint8  `validate:"gte=1,lte=10"`
    Sex      string `validate:"oneof=female male"`
}

func main() {
    validate := validator.New()

    user1 := User{Username: "asong", Age: 11, Sex: "null"}
    err := validate.Struct(user1)
    if err != nil {
        fmt.Println(err)
    }

    user2 := User{Username: "asong111", Age: 8, Sex: "male"}
    err = validate.Struct(user2)
    if err != nil {
        fmt.Println(err)
    }

}
```

我们在结构体定义validator标签的tag，使用`validator.New()`创建一个验证器，这个验证器可以指定选项、添加自定义约束，然后在调用他的`Struct()`方法来验证各种结构对象的字段是否符合定义的约束。

上面的例子，我们在User结构体中，有三个字段：

- Name：通过min和max来进行约束，Name的字符串长度为[6,10]之间。
- Age：通过gte和lte对年轻的范围进行约束，age的大小大于1，小于10。
- Sex：通过oneof对值进行约束，只能是所列举的值，oneof列举出性别为男士🚹和女士🚺(不是硬性规定奥，可能还有别的性别)。

所以`user1`会进行报错，错误信息如下：

```
Key: 'User.Name' Error:Field validation for 'Name' failed on the 'min' tag
Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag
Key: 'User.Sex' Error:Field validation for 'Sex' failed on the 'oneof' tag
```

### 一些约束

#### 字符串约束

* `excludesall`：不包含参数中任意的 UNICODE 字符，例如`excludesall=ab`；
* `excludesrune`：不包含参数表示的 rune 字符，`excludesrune=asong`；
* `startswith`：以参数子串为前缀，例如`startswith=hi`；
* `endswith`：以参数子串为后缀，例如`endswith=bye`。
* `contains=`：包含参数子串，例如`contains=email`；
* `containsany`：包含参数中任意的 UNICODE 字符，例如`containsany=ab`；
* `containsrune`：包含参数表示的 rune 字符，例如`containsrune=asong；
* `excludes`：不包含参数子串，例如`excludes=email`；

#### 范围约束

范围约束的字段类型分为三种：

- 对于数值，我们则可以约束其值
- 对于切片、数组和map，我们则可以约束其长度
- 对于字符串，我们则可以约束其长度

常用tag介绍：

- `ne`：不等于参数值，例如`ne=5`；
- `gt`：大于参数值，例如`gt=5`；
- `gte`：大于等于参数值，例如`gte=50`；
- `lt`：小于参数值，例如`lt=50`；
- `lte`：小于等于参数值，例如`lte=50`；
- `oneof`：只能是列举出的值其中一个，这些值必须是数值或字符串，以空格分隔，如果字符串中有空格，将字符串用单引号包围，例如`oneof=male female`。
- `eq`：等于参数值，注意与`len`不同。对于字符串，`eq`约束字符串本身的值，而`len`约束字符串长度。例如`eq=10`；
- `len`：等于参数值，例如`len=10`；
- `max`：小于等于参数值，例如`max=10`；
- `min`：大于等于参数值，例如`min=10`

#### Fields约束

- `eqfield`：定义字段间的相等约束，用于约束同一结构体中的字段。例如：`eqfield=Password`
- `eqcsfield`：约束统一结构体中字段等于另一个字段（相对），确认密码时可以使用，例如：`eqfiel=ConfirmPassword`
- `nefield`：用来约束两个字段是否相同，确认两种颜色是否一致时可以使用，例如：`nefield=Color1`
- `necsfield`：约束两个字段是否相同（相对）

#### 常用约束

- `unique`：指定唯一性约束，不同类型处理不同：
  - 对于map，unique约束没有重复的值
  - 对于数组和切片，unique没有重复的值
  - 对于元素类型为结构体的碎片，unique约束结构体对象的某个字段不重复，使用`unique=field`指定字段名
- `email`：使用`email`来限制字段必须是邮件形式，直接写`eamil`即可，无需加任何指定。
- `omitempty`：字段未设置，则忽略
- `-`：跳过该字段，不检验；
- `|`：使用多个约束，只需要满足其中一个，例如`rgb|rgba`；
- `required`：字段必须设置，不能为默认值；