import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';

import { NzModalService } from 'ng-zorro-antd/modal';
import { NzNotificationService } from 'ng-zorro-antd/notification';

import { FunctionsService } from 'src/app/services/functions.service';

@Component({
  selector: 'app-new-account',
  templateUrl: './new-account.component.html',
  styleUrls: ['./new-account.component.less'],
})
export class NewAccountComponent implements OnInit {
  constructor(
    private modal: NzModalService,
    private functionsService: FunctionsService,
    private notify: NzNotificationService
  ) {}

  loadingButton = false;

  public newAccountForm = new FormGroup({
    nickname: new FormControl(),
    publicKey: new FormControl(),
  });

  public closeModal() {
    this.modal.closeAll();
  }

  public createAccount() {
    const payload = {
      userId: localStorage.getItem('id')!,
      publicKey: this.newAccountForm.value.publicKey,
      nickname: this.newAccountForm.value.nickname,
    };

    this.functionsService.createAccount(payload).subscribe(
      (response) => {
        this.notify.success('Success', 'Account created');
        this.closeModal();
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  ngOnInit(): void {}
}
