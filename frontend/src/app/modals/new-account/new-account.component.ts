import { Component, OnInit } from '@angular/core';
import { NzModalService } from 'ng-zorro-antd/modal';

@Component({
  selector: 'app-new-account',
  templateUrl: './new-account.component.html',
  styleUrls: ['./new-account.component.less'],
})
export class NewAccountComponent implements OnInit {
  constructor(private modal: NzModalService) {}

  public closeModal() {
    this.modal.closeAll();
  }

  public createAccount() {
    console.log('create account');
  }

  ngOnInit(): void {}
}
