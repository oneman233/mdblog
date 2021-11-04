import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'],
})
export class NavbarComponent implements OnInit {

  constructor(private location: Location) { }

  ngOnInit(): void {
  }

  // 不处于根目录则展示返回按钮
  get isRoot(): boolean {
    return this.location.isCurrentPathEqualTo("/articles");
  }

  goBack(): void {
    this.location.back();
  }
}
