import { Component, OnInit, TemplateRef, Type } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { NewPaymentLinksComponent } from 'src/app/modals/new-payment-link/new-payment-link.component';
import { QrcodeLinkComponent } from 'src/app/modals/qrcode-link/qrcode-link.component';
import {
  FunctionsService,
  FunctionsServiceInterface,
} from 'src/app/services/functions.service';

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
  selector: 'app-payment-links',
  templateUrl: './payment-links.component.html',
  styleUrls: ['./payment-links.component.less'],
})
export class PaymentLinksComponent implements OnInit {
  constructor(
    private functionsService: FunctionsService,
    private notify: NzNotificationService,
    private modal: NzModalService
  ) {}

  public links!: FunctionsServiceInterface.Receive.listLink[];

  public infosLink!: InfosLink;
  private modalQrCode!: NzModalRef;

  public listLink() {
    this.functionsService.listLink().subscribe(
      (response) => {
        response.forEach((link) => {
          link.expand = false;
        });

        this.links = response;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public newLink() {
    this.modal
      .create({
        nzTitle: 'New Payment Link',
        nzContent: NewPaymentLinksComponent,
        nzFooter: null,
        nzWidth: '50%',
      })
      .afterClose.subscribe(() => {
        this.listLink();
      });
  }

  private openModal(
    nzContent: string | TemplateRef<{}> | Type<unknown> | undefined,
    infosLink?: InfosLink
  ): void {
    this.modalQrCode = this.modal.create({
      nzContent,
      nzFooter: null,
      nzWidth: '43.188rem',
      nzMaskClosable: false,
      nzBodyStyle: {
        padding: '40px 60px',
      },
      nzComponentParams: {
        infosLink,
      },
    });
  }

  public generateQRCode(link: FunctionsServiceInterface.Receive.listLink) {
    this.openModal(QrcodeLinkComponent, link);
  }

  ngOnInit(): void {
    this.listLink();
  }
}
