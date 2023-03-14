import { Component, OnInit } from '@angular/core';
import { NzModalRef } from 'ng-zorro-antd/modal';

import * as qrcode from 'qrcode-generator';

interface InfosLink {
  ID: string;
  Nickname: string;
  UserId: string;
  AccountId: string;
  Link: string;
  Reference: string;
  Recipient: string;
  Network: string;
  ExpectedAmount: number;
  AmountReceived: number;
  Status: string;
  CreatedAt: string;
  ReceivedAt: string;
}

@Component({
  selector: 'app-qrcode-link',
  templateUrl: './qrcode-link.component.html',
  styleUrls: ['./qrcode-link.component.less'],
})
export class QrcodeLinkComponent implements OnInit {
  constructor(private modal: NzModalRef) {}

  private infosLink!: InfosLink;
  public linkQrCode!: string;

  public generateQRCode(text: string): void {
    this.linkQrCode = text;
    const qr = qrcode(0, 'Q');
    qr.addData(text);
    qr.make();
    const qrCode = document.getElementById('qrCode');
    qrCode!.innerHTML = qr.createImgTag(4, 0);
  }

  ngOnInit() {
    this.generateQRCode(this.infosLink.Link);
  }
}
