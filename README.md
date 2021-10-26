# rumvalidate-go

基于函数式编程实现的参数校验器：轻量，高性能，自由可扩展

## 性能测试

和go-playground/validator进行和简单的对比测试

示例代码

    // rumis/govalidate
    func BenchmarkRumValidate(b *testing.B) {
        params := map[string]interface{}{
            "curpage": 2,
            "perpage": 14,
            "u":       "http://www.baidu.com",
            "range":   2,
        }
        rules := map[string]R.FilterItem{
            "curpage": govalidate.Filter([]R.Validator{R.Required()}),
            "perpage": govalidate.Filter([]R.Validator{R.Required()}),
            "u":       govalidate.Filter([]R.Validator{R.Required(), R.Url()}),
            "range":   govalidate.Filter([]R.Validator{R.Required(), R.Between(0, 100)}),
        }
        for i := 0; i < b.N; i++ {
            _, _, _ = govalidate.Validate(params, rules)
        }
    }

    // go-playground/validator/v10
    func BenchmarkValidate(b *testing.B) {
        p := Paginate{
            Perpage: 12,
            Curpage: 1,
            U:       "http://www.baidu.com",
            Range:   22,
        }
        for i := 0; i < b.N; i++ {
            _ = validator.New().Struct(p)
        }
    }

测试结果

    go test -benchmem -bench=.
    goos: linux
    goarch: amd64
    pkg: liumurong.org/debug/validatet
    cpu: Intel(R) Xeon(R) CPU E5-26xx v4
    BenchmarkRumValidate-2            785211              1589 ns/op             696 B/op          8 allocs/op
    BenchmarkValidate-2                26142             45706 ns/op           14959 B/op        188 allocs/op
    PASS
    ok      liumurong.org/debug/validatet   2.961s 

内存占用测试:

GODEBUG='gctrace=1' go test -benchmem -bench=BenchmarkRumValidate

    cpu: Intel(R) Xeon(R) CPU E5-26xx v4
    BenchmarkRumValidate-2          
    gc 3 @0.012s 1%: 0.040+0.90+0.003 ms clock, 0.080+0/0.078/0.77+0.006 ms cpu, 1->1->1 MB, 4 MB goal, 2 P (forced)
    gc 4 @0.014s 2%: 0.035+0.51+0.003 ms clock, 0.071+0/0.11/0.85+0.007 ms cpu, 1->1->1 MB, 4 MB goal, 2 P (forced)
    gc 13 @0.087s 4%: 0.024+1.0+0.026 ms clock, 0.048+0.18/0.075/0.89+0.053 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 14 @0.094s 5%: 0.022+4.4+0.030 ms clock, 0.045+0/1.2/4.1+0.061 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 21 @0.156s 5%: 0.031+1.9+0.028 ms clock, 0.062+0/1.0/1.2+0.057 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 22 @0.164s 5%: 0.036+2.3+0.029 ms clock, 0.072+0/1.3/1.4+0.059 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 31 @0.242s 5%: 0.058+3.1+0.062 ms clock, 0.11+0/1.1/1.3+0.12 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 32 @0.250s 5%: 0.027+1.8+0.025 ms clock, 0.054+0/0.96/1.1+0.051 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 33 @0.259s 5%: 0.025+1.4+0.026 ms clock, 0.050+0.35/0.26/1.0+0.053 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 40 @0.316s 5%: 0.023+1.5+0.027 ms clock, 0.047+0.13/0.76/0.50+0.054 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 41 @0.325s 5%: 0.025+7.1+0.027 ms clock, 0.051+0/2.3/5.2+0.055 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 51 @0.420s 5%: 0.023+4.1+0.028 ms clock, 0.046+0/1.2/3.2+0.056 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 52 @0.430s 5%: 0.018+2.2+0.023 ms clock, 0.037+0/0.89/1.4+0.047 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 62 @0.528s 5%: 0.025+6.7+0.030 ms clock, 0.050+0/0.88/0.80+0.060 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 63 @0.543s 5%: 0.023+3.7+0.016 ms clock, 0.047+0.79/0.49/0+0.032 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 76 @0.773s 7%: 0.034+10+0.025 ms clock, 0.069+0.18/0.12/0.99+0.050 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 77 @0.799s 7%: 0.036+4.1+0.004 ms clock, 0.073+4.0/0.066/1.2+0.009 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 78 @0.812s 7%: 0.031+3.4+0.021 ms clock, 0.063+0.21/0.12/0.94+0.043 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 92 @0.960s 7%: 0.035+1.4+0.029 ms clock, 0.070+0.20/0.087/1.1+0.059 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 93 @0.972s 7%: 0.025+1.2+0.026 ms clock, 0.051+0.18/0.12/0.88+0.052 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 105 @1.090s 6%: 0.020+1.1+0.022 ms clock, 0.041+0.15/0.15/0.90+0.044 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 106 @1.102s 6%: 0.043+1.5+0.025 ms clock, 0.086+0.32/0.081/1.2+0.051 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 117 @1.212s 6%: 0.020+5.6+0.027 ms clock, 0.040+0/2.3/4.0+0.055 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 118 @1.222s 6%: 0.020+0.95+0.026 ms clock, 0.040+0.18/0.061/0.80+0.053 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 132 @1.364s 6%: 0.093+3.9+0.031 ms clock, 0.18+0/0.97/2.9+0.063 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 133 @1.373s 6%: 0.019+4.2+0.027 ms clock, 0.039+0/1.2/3.9+0.054 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 146 @1.510s 6%: 0.033+3.3+0.004 ms clock, 0.066+0.98/1.0/0+0.008 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 147 @1.522s 7%: 13+5.3+0.010 ms clock, 26+0/6.8/0+0.020 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    gc 148 @1.549s 7%: 0.54+1.4+0.024 ms clock, 1.0+0.18/0.45/0.95+0.049 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
    605376              2518 ns/op             696 B/op          8 allocs/op
    PASS
    ok      liumurong.org/debug/validatet   1.568s


GODEBUG='gctrace=1' go test -benchmem -bench=BenchmarkValidate

    cpu: Intel(R) Xeon(R) CPU E5-26xx v4
    BenchmarkValidate-2     
    gc 3 @0.017s 6%: 0.045+0.63+0.003 ms clock, 0.090+0/0.19/0.98+0.007 ms cpu, 1->1->1 MB, 4 MB goal, 2 P (forced)
    gc 4 @0.029s 4%: 0.049+1.4+0.005 ms clock, 0.099+0/0.16/2.5+0.010 ms cpu, 2->2->2 MB, 4 MB goal, 2 P (forced)
    gc 9 @0.158s 5%: 0.092+10+0.035 ms clock, 0.18+2.3/4.1/5.8+0.070 ms cpu, 17->19->11 MB, 18 MB goal, 2 P
    gc 10 @0.195s 5%: 0.082+11+0.056 ms clock, 0.16+3.8/0.61/10+0.11 ms cpu, 21->22->13 MB, 22 MB goal, 2 P
    gc 11 @0.235s 6%: 0.095+15+0.026 ms clock, 0.19+1.5/5.0/11+0.053 ms cpu, 25->26->14 MB, 26 MB goal, 2 P
    gc 16 @0.608s 7%: 0.12+14+0.005 ms clock, 0.25+0/3.0/23+0.010 ms cpu, 41->41->21 MB, 45 MB goal, 2 P (forced)
    gc 17 @0.671s 7%: 0.13+14+0.024 ms clock, 0.26+1.9/8.0/10+0.049 ms cpu, 40->42->20 MB, 43 MB goal, 2 P
    gc 18 @0.728s 7%: 0.12+15+0.026 ms clock, 0.24+3.8/5.7/12+0.052 ms cpu, 38->40->21 MB, 40 MB goal, 2 P
    gc 24 @1.224s 7%: 0.16+28+0.032 ms clock, 0.32+1.8/12/22+0.065 ms cpu, 56->60->32 MB, 60 MB goal, 2 P
    gc 25 @1.329s 8%: 0.48+31+0.032 ms clock, 0.97+2.6/16/23+0.064 ms cpu, 61->64->34 MB, 65 MB goal, 2 P
    gc 28 @1.605s 8%: 0.17+5.1+0.035 ms clock, 0.34+0/0.16/10+0.071 ms cpu, 44->44->9 MB, 78 MB goal, 2 P (forced)
    gc 29 @1.639s 8%: 0.069+7.6+0.083 ms clock, 0.13+2.4/2.5/6.0+0.16 ms cpu, 17->18->9 MB, 18 MB goal, 2 P
    gc 30 @1.668s 8%: 0.073+8.8+0.030 ms clock, 0.14+2.0/3.1/5.7+0.061 ms cpu, 18->19->11 MB, 19 MB goal, 2 P
    gc 37 @2.035s 7%: 0.16+20+0.027 ms clock, 0.32+3.2/9.9/14+0.055 ms cpu, 47->50->27 MB, 51 MB goal, 2 P
    gc 38 @2.111s 7%: 0.16+22+0.028 ms clock, 0.32+6.0/5.9/19+0.056 ms cpu, 50->54->29 MB, 55 MB goal, 2 P
    gc 43 @2.593s 8%: 0.21+61+0.30 ms clock, 0.42+9.9/33/42+0.60 ms cpu, 67->71->37 MB, 72 MB goal, 2 P
    gc 44 @2.754s 8%: 0.20+28+0.027 ms clock, 0.40+6.5/14/20+0.054 ms cpu, 70->73->38 MB, 74 MB goal, 2 P
    25482             45934 ns/op           14962 B/op        188 allocs/op
    PASS
    ok      liumurong.org/debug/validatet   2.806s