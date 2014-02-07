def waysToMakeChange(amount, coins, waysFound = 0):
  if (amount==0 or len(coins)<=1):
    return (amount==0) or (amount % coins[0]==0)
  branches = 0
  for i in range(0,1+amount/coins[0]):
    branches+=waysToMakeChange(amount-i*coins[0],coins[1:])
  return branches
 
coins=[200,100,50,20,10,5,2,1]
print(waysToMakeChange(200,coins))