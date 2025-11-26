import nodemailer from 'nodemailer';
import { Injectable, OnModuleInit } from '@nestjs/common';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@auth/config';

@Injectable()
export class MailService implements OnModuleInit {
  private transporter: nodemailer.Transporter;

  constructor(private readonly configService: ConfigService<Config>) {}

  onModuleInit() {
    this.transporter = nodemailer.createTransport({
      service: 'gmail',
      auth: {
        user: this.configService.get('mail.auth.user'),
        pass: this.configService.get('mail.auth.pass'),
      },
    });
  }

  async sendMail({
    to,
    subject,
    html,
  }: {
    to: string;
    subject: string;
    html: string;
  }) {
    try {
      return await this.transporter.sendMail({
        from: this.configService.get('mail.auth.user'),
        to,
        subject,
        html,
      });
    } catch (error) {
      throw new Error(`Error sending email: ${error}`);
    }
  }
}
