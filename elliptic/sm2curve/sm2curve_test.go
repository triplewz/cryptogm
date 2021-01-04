package sm2curve

import (
	"math/big"
	"testing"
)

func TestP256Point(t *testing.T)  {
	TestP256_Point()
}

func TestP256InternalFunc(t *testing.T)  {
	Test_p256InternalFunc()
}

func TestP256Func(t *testing.T)  {
	Test_p256Func()
}

func TestArch64(t *testing.T)  {
	Test_arch64()
}

//func TestADD(t *testing.T) {
//	ADD()
//}
//
//func TestPointAdd(t *testing.T) {
//	PointAdd()
//}

func TestBaseMul(t *testing.T)  {
	//c := P256()
	//k,_ := randFieldElement(c,rand.Reader)
	//fmt.Printf("%x\n",k)
	k,_ := new(big.Int).SetString("9d0fe3e894232063af63863b612f280e3bbd5598d70400de6fb4443eb0fbc52b",16)

	k1 := make([]uint64, 4)
	fromBig(k1, k)

	p := new(p256Point)

	p.p256BaseMult(k1)

	x, y := Uint64ToAffine(p.xyz[:])
	Hexprint(x.Bytes())
	Hexprint(y.Bytes())
}