import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';

import { NzNotificationService } from 'ng-zorro-antd/notification';

import * as qrcode from 'qrcode-generator';
import { PermalinkService } from 'src/app/services/permalink.service';

@Component({
  selector: 'app-create-permalink',
  templateUrl: './create-permalink.component.html',
  styleUrls: ['./create-permalink.component.less'],
})
export class CreatePermalinkComponent implements OnInit {
  constructor(
    private permalinkService: PermalinkService,
    private notify: NzNotificationService
  ) {}

  public showQrCode = false;

  public linkQrCode!: string;
  public codePermalink!: string;

  public userId = window.location.href.split('/')[4];

  public createPermalinkForm = new FormGroup({
    expectedAmount: new FormControl(),
  });

  public generateQRCode(text: string): void {
    this.linkQrCode = text;
    const qr = qrcode(0, 'Q');
    qr.addData(text);
    qr.make();
    const qrCode = document.getElementById('qrCode');
    qrCode!.innerHTML = qr.createImgTag(4, 0);
  }

  public createPermalink() {
    this.showQrCode = true;

    const payload = {
      userId: this.userId,
      expectedAmount: this.createPermalinkForm.value.expectedAmount,
    };

    this.permalinkService.createPermalink(payload).subscribe(
      (response) => {
        this.codePermalink = response.code;
        this.generateQRCode(response.link);
        this.showQrCode = true;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public backPage() {
    this.showQrCode = false;
  }

  ngOnInit(): void {}
}
