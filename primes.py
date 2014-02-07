import math
import sys

def primes(x):
  if x<2:
    print("Primes start at +2, do it right")
    return
  primes = [True]*x
  sqrt = int(math.ceil(x**.5))
  for i in range(2, sqrt):
    if (primes[i]):
      for j in range(i*i, x, i):
        primes[j] = False
  c = 0
  for i in range(2, x):
    if (primes[i]):
      c += 1
  print(c)

primes(int(sys.argv[1]))