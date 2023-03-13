import { Component, OnInit } from '@angular/core';

import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less'],
})
export class LoginComponent implements OnInit {
  constructor(private notify: NzNotificationService) {}

  public passwordVisible = false;

  loadingButton = false;

  public login() {
    console.log('Login');
  }

  ngOnInit(): void {}
}
