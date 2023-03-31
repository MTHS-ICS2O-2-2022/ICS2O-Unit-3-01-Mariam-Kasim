
// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mariam Kasim
// Created on: Mar 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates area of a Trapezoid
 */
function calculate () {
  // input
  const baseA = parseInt(document.getElementById('base-a').value)
  const baseB = parseInt(document.getElementById('base-b').value)
  const height = parseInt(document.getElementById('height').value)



  // process
  const area = [(baseA + baseB) / 2] * height

  // output
  document.getElementById('area').innerHTML = 'area is: ' + area + ' cm²'
}
