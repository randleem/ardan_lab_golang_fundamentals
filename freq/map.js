let myArr = [1, 2, 3];

let plusArr = myArr.map((number) => {
  return number + 1;
});

//plusArr => [2,3,4]

let orderIds = [{ id: 123 }, { id: 456 }, { id: 789 }];

let orderIds2 = orderIds.map((orderid) => {
  return orderid.id;
});
//orderIds2 = [123, 456,789]
console.log(orderIds2);
