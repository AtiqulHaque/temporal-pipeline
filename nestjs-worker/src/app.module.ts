import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { InvoiceModule } from './invoice/invoice.module';
import { Invoice } from './invoice/invoice.entity';
import { TemporalModule } from './temporal/temporal.module';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: process.env.MYSQL_HOST ?? 'localhost',
      port: Number(process.env.MYSQL_PORT ?? 3306),
      username: process.env.MYSQL_USER ?? 'invoice',
      password: process.env.MYSQL_PASSWORD ?? 'invoice',
      database: process.env.MYSQL_DATABASE ?? 'invoices',
      entities: [Invoice],
      synchronize: process.env.MYSQL_SYNCHRONIZE !== 'false',
      logging: process.env.MYSQL_LOGGING === 'true',
    }),
    InvoiceModule,
    TemporalModule,
  ],
})
export class AppModule {}
