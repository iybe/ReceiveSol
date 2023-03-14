import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';
import { FormGroup, FormControl } from '@angular/forms';

import { NzNotificationService } from 'ng-zorro-antd/notification';
import { SsoService } from 'src/app/services/sso.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less'],
})
export class LoginComponent implements OnInit {
  constructor(
    private router: Router,
    private ssoService: SsoService,
    private notify: NzNotificationService
  ) {}

  public passwordVisible = false;

  loadingButton = false;

  public loginForm = new FormGroup({
    username: new FormControl(),
    password: new FormControl(),
  });

  public clearLocalStorage() {
    localStorage.clear();
  }

  public newRegister() {
    this.router.navigate(['/register']);
  }

  public login() {
    this.loadingButton = true;

    const payload = {
      username: this.loginForm.value.username,
      password: this.loginForm.value.password,
    };

    this.ssoService.login(payload).subscribe(
      (response) => {
        localStorage.setItem('token', response.token);
        localStorage.setItem('id', response.id);
        localStorage.setItem('username', response.username);

        this.notify.success('Success', 'Login success');
        this.router.navigate(['/receivesol']);
        this.loadingButton = false;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
        this.loadingButton = false;
      }
    );
  }

  ngOnInit(): void {
    this.clearLocalStorage();
  }
}
