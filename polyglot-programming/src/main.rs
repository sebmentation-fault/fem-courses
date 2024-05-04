fn main() {

  let mut a = vec![];
  let mut b = a;

  b.push(1);
  a.push(1);

  println!("{:?}", b);

}
