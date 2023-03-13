import { Component, OnInit } from '@angular/core';
import { NzModalService } from 'ng-zorro-antd/modal';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { NewAccountComponent } from 'src/app/modals/new-account/new-account.component';
import {
  FunctionsService,
  FunctionsServiceInterface,
} from 'src/app/services/functions.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.less'],
})
export class AccountComponent implements OnInit {
  constructor(
    private functionsService: FunctionsService,
    private notify: NzNotificationService,
    private modal: NzModalService
  ) {}

  public accounts!: FunctionsServiceInterface.Receive.listAccount[];

  public listAccount() {
    this.functionsService.listAccount().subscribe(
      (response) => {
        this.accounts = response;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public newAccount() {
    this.modal
      .create({
        nzTitle: 'New Account',
        nzContent: NewAccountComponent,
        nzFooter: null,
        nzWidth: '50%',
      })
      .afterClose.subscribe(() => {
        this.listAccount();
      });
  }

  ngOnInit(): void {
    this.listAccount();
  }
}
