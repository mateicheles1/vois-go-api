"use strict";

class Person {
  constructor(name, address, phone, course) {
    this.name = name;
    this.address = address;
    this.phone = phone;
    this.course = course;
  }

  getInfo() {
    console.log(
      `Name: ${this.name}, Address: ${this.address}, Phone: ${this.phone}, Course: ${this.course}`
    );
  }
}

const firstPerson = new Person(
  "andreea",
  "iuliu maniu 15h",
  07777777,
  "Frontend web dev"
);

firstPerson.getInfo();
