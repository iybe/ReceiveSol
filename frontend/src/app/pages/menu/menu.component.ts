import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.less'],
})
export class MenuComponent implements OnInit {
  constructor(private router: Router) {}

  public userName: string = localStorage.getItem('username')!;

  public account = true;
  public paymantLinks = false;
  public permalinks = false;

  public changePage(page: string) {
    if (page === 'Account') {
      this.account = true;
      this.paymantLinks = false;
      this.permalinks = false;
    }

    if (page === 'paymantLinks') {
      this.paymantLinks = true;
      this.account = false;
      this.permalinks = false;
    }

    if (page === 'permalinks') {
      this.permalinks = true;
      this.paymantLinks = false;
      this.account = false;
    }
  }

  public verifyToken() {
    if (localStorage.getItem('token') == null) {
      this.router.navigate(['/login']);
    }
  }

  public singOut() {
    localStorage.clear();
    this.router.navigate(['/login']);
  }

  ngOnInit(): void {
    this.verifyToken();
  }
}
