# Euler's totient function

# Links

* [Euler's totient function - Wikipedia](https://en.wikipedia.org/wiki/Euler's_totient_function)

# Euler's product formula

![euler_product_formula](img/euler_product_formula.svg)

![epf](img/epf.svg)

where the product is over the distinct prime numbers dividing n.

The proof of Euler's product formula depends on two important facts. 

**The function is multiplicative**

This means that if gcd(m, n) = 1, then φ(mn) = φ(m) φ(n). 

**Value for a prime power argument**

If p is prime and k ≥ 1, then 

![value](img/value_for_a_prime_power_argument.svg)


Proof: since p is a prime number the only possible values of gcd(p^k, m) are 1,
p, p^2, ..., p^k, and the only way for gcd(p^k, m) to not equal 1 is for m to be
a multiple of p. The multiples of p that are less than or equal to p^k are p,
2p, 3p, ..., p^(k − 1)p = p^k, and there are p^(k − 1) of them. Therefore, the other
p^k − p^(k − 1) numbers are all relatively prime to p^k. 

上面这个很有意思，为了找得gcd(p^k,m)=1，m的个数。转换成去找gcd(p^k,m) != 1,m的个数
，这时m MUST= x * p，这个个数有p^(k-1)个, p^k - p^(k-1)就是俺要的gcd(p^k,m)=1的
m的个数

**Example**

![alt](img/epf.eg.svg)

In words, this says that the distinct prime factors of 36 are 2 and 3; half of
the thirty-six integers from 1 to 36 are divisible by 2, leaving eighteen;
a third of those are divisible by 3, leaving twelve numbers that are coprime to 36.
And indeed there are twelve positive integers that are coprime with 36 and
lower than 36: 1, 5, 7, 11, 13, 17, 19, 23, 25, 29, 31, and 35. 


