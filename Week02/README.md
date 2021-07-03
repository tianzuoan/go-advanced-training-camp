学习笔记

作业： 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：在dao层中遇到sql.ErrNoRows时都应该使用warp进行包装，再由业务层进行判断，将决定权放在业务层，业务层决定是中断业务还是进行降级处理，遵循错误只处理一次的原则

# 异常处理

## Error vs Exception

### errors.New()返回的是指针

### type error是一个接口，只有一个Error()方法

### go使用多参数返回来支持返回error

### go使用panic机制来处理意外情况，一般的异常错误使用error

### Error are values

### 某个goroutine的panic只能由其当前goroutine的defer逻辑捕获，其他goroutine无法处理

## Error Type

### sentinel error 预定义的错误

- 成为API的公共部分
- 在两个包之间创建了依赖
- 依赖Error()方法的输出，并且不好扩展错误信息。但该方法一般给程序员调试使用
- 结论：尽量避免使用sentinel error

### Error types

### Opaque errors

## GO2 Error Inspection

## Handing Error

*XMind - Trial Version*