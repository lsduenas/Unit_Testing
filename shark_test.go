package hunt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// Arrange - Given
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  26,
	}
	p := Prey{
		name:  "pcsito",
		speed: 20,
	}

	// Act - When
	errObtained := s.Hunt(&p)

	// Assert - Then
	assert.Nil(t, errObtained)
	assert.Equal(t, false, s.hungry) // not hungry
	assert.Equal(t, true, s.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  true,
		speed:  26,
	}
	p := Prey{
		name:  "pcsito",
		speed: 20,
	}
	var err error = errors.New("cannot hunt, i am really tired")

	// act
	errObtained := s.Hunt(&p)

	// assert
	assert.EqualErrorf(t, errObtained, err.Error(), "")
}

func TestSharkCannotHuntBecauseIsNotHungry(t *testing.T) {
	// arrange
	s := Shark{
		hungry: false,
		tired:  false,
		speed:  26,
	}
	p := Prey{
		name:  "pcsito",
		speed: 20,
	}
	var err error = errors.New("cannot hunt, i am not hungry")

	// act
	errObtained := s.Hunt(&p)

	// assert
	assert.EqualErrorf(t, errObtained, err.Error(), "Test executed successfully")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  16,
	}
	p := Prey{
		name:  "pcsito",
		speed: 20,
	}
	var err error = errors.New("could not catch it")

	// act
	errObtained := s.Hunt(&p)

	// assert
	assert.EqualErrorf(t, errObtained, err.Error(), "Test executed successfully")
}

func TestSharkHuntNilPrey(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  20,
	}
	var p *Prey
	var err error = errors.New("nil prey")

	// act
	errObtained := s.Hunt(p)

	// assert
	assert.EqualErrorf(t, errObtained, err.Error(), "Test executed successfully")
}
