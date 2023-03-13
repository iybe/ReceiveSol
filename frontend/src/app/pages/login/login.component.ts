import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';
import { FormGroup, FormControl } from '@angular/forms';

import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less'],
})
export class LoginComponent implements OnInit {
  constructor(private router: Router, private notify: NzNotificationService) {}

  public passwordVisible = false;

  loadingButton = false;

  public loginForm = new FormGroup({
    username: new FormControl(),
    password: new FormControl(),
  });

  public newRegister() {
    this.router.navigate(['/register']);
  }

  public login() {
    console.log(this.loginForm.value);
  }

  ngOnInit(): void {}
}
