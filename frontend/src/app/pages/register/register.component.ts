import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

import { NzNotificationService } from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.less'],
})
export class RegisterComponent implements OnInit {
  constructor(private router: Router, private notify: NzNotificationService) {}

  public passwordVisible = false;

  loadingButton = false;

  public registerForm = new FormGroup({
    username: new FormControl(),
    password: new FormControl(),
    passwordConfirm: new FormControl(),
  });

  public backToLogin() {
    this.router.navigate(['/login']);
  }

  public registerUser() {
    console.log(this.registerForm.value);
  }

  ngOnInit() {}
}
