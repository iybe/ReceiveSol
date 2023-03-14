import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { NzModalService } from 'ng-zorro-antd/modal';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import {
  FunctionsService,
  FunctionsServiceInterface,
} from 'src/app/services/functions.service';

@Component({
  selector: 'app-new-payment-links',
  templateUrl: './new-payment-link.component.html',
  styleUrls: ['./new-payment-link.component.less'],
})
export class NewPaymentLinksComponent implements OnInit {
  constructor(
    private modal: NzModalService,
    private functionsService: FunctionsService,
    private notify: NzNotificationService
  ) {}

  public accounts!: FunctionsServiceInterface.Receive.listAccount[];

  public networks = ['testnet', 'devnet', 'mainnet'];

  public newPaymentLinkForm = new FormGroup({
    nickname: new FormControl(),
    network: new FormControl(),
    recipient: new FormControl(),
    expectedAmount: new FormControl(),
  });

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

  public closeModal() {
    this.modal.closeAll();
  }

  public createPaymentLink() {
    const payload = {
      userId: localStorage.getItem('id')!,
      nickname: this.newPaymentLinkForm.value.nickname,
      recipient: this.newPaymentLinkForm.value.recipient,
      network: this.newPaymentLinkForm.value.network,
      expectedAmount: this.newPaymentLinkForm.value.expectedAmount,
    };

    this.functionsService.createLink(payload).subscribe(
      (response) => {
        this.notify.success('Success', 'Account created');
        this.closeModal();
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  ngOnInit(): void {
    this.listAccount();
  }
}
