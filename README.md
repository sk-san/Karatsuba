<h1>Karatsuba</h1>
Karatsuba multiplication is one of famous algorithm to perform Polynomial and fast multiplication developed by Anatoly Alexeyevich Karatsuba(Анато́лий Алексе́евич Карацу́ба). This algorithm is renowned for the typical divide-and-conquer algorithm that reduce the number of multiplication but increase the number of additions because of the cost of additions is fewer than multiplications. As a result, the whole calculating cost is almost 3/4 compared to the case without reducing the number of multiplication.

<h1>How it behave</h1>
A = 1234
B = 5678

a = 12
b = 34
c = 56
d = 78
m = 4 (length of A or B)

<h2>Step1</h2>
Compute   a x c = 672
<h2>Step2</h2>
Compute b x d = 2652
<h2>Step3</h2>
Compute (a + b)(c + d) = 134 * 46 = 6264
<h2>Step4</h2>
Compute Step3 - (Step1 + Step2) = 2840
<h2>Step5</h2>
Compute Step1 ^ m + Step2 + Step4 ^ (m/2)
