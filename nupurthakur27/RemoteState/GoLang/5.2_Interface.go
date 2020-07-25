package main

type Cuboid interface {
	BoxSurfaceArea() float64
	Volume() float64
}

type values struct {
	breadth float64
	length float64
	height float64
}

func (v values) BoxSurfaceArea() float64{
	return 2*(v.breadth*v.length +v.length*v.height+v.height*v.breadth)
}
func (v values) Volume() float64  {
	return v.length*v.breadth*v.height
}