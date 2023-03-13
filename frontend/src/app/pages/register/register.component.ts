import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

import { NzNotificationService } from 'ng-zorro-antd/notification';
import { SsoService } from 'src/app/services/sso.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.less'],
})
export class RegisterComponent implements OnInit {
  constructor(
    private router: Router,
    private ssoService: SsoService,
    private notify: NzNotificationService
  ) {}

  public passwordVisible = false;
  public passwordConfirmVisible = false;

  loadingButton = false;

  public registerForm = new FormGroup({
    username: new FormControl(),
    password: new FormControl(),
    passwordConfirm: new FormControl(),
  });

  public backToLogin() {
    this.router.navigate(['/login']);
  }

  public passwordsMatch() {
    return (
      this.registerForm.value.password !==
      this.registerForm.value.passwordConfirm
    );
  }

  public registerUser() {
    this.loadingButton = true;

    const payload = {
      username: this.registerForm.value.username,
      password: this.registerForm.value.password,
    };

    this.ssoService.register(payload).subscribe(
      (response) => {
        this.notify.success('Success', 'Register success');
        this.router.navigate(['/login']);
        this.loadingButton = false;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
        this.loadingButton = false;
      }
    );
  }

  ngOnInit() {}
}
