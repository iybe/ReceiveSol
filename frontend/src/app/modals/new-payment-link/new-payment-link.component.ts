import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';

import { NzModalService } from 'ng-zorro-antd/modal';
import { NzNotificationService } from 'ng-zorro-antd/notification';

import * as qrcode from 'qrcode-generator';

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
  public linkQrCode!: string;

  public networks = ['testnet', 'devnet', 'mainnet'];

  public showQrCode = false;

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

  public generateQRCode(text: string): void {
    this.linkQrCode = text;
    const qr = qrcode(0, 'Q');
    qr.addData(text);
    qr.make();
    const qrCode = document.getElementById('qrCode');
    qrCode!.innerHTML = qr.createImgTag(4, 0);
  }

  public createPaymentLink() {
    this.showQrCode = true;

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
        this.showQrCode = true;
        this.generateQRCode(response.link);
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public backPage() {
    this.showQrCode = false;
  }

  ngOnInit(): void {
    this.listAccount();
  }
}
