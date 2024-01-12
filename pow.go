func myPow(x float64, n int) float64 {
    res := 1.0;
    var nums uint32 = uint32(n);
    if n < 0{
        x = 1/x;
        nums = -nums;
    }
    for nums!=0{
        if nums & 1 == 1{
            res*=x
        }
        x*=x
        nums = nums >> 1
    }
    return res
}
