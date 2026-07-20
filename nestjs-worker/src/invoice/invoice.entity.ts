import {
  Column,
  CreateDateColumn,
  Entity,
  PrimaryGeneratedColumn,
} from 'typeorm';

export type InvoiceStatus = 'pending' | 'created' | 'failed';

@Entity('invoices')
export class Invoice {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Column({ name: 'invoice_number', length: 64, unique: true })
  invoiceNumber!: string;

  @Column({ name: 'customer_name', length: 255 })
  customerName!: string;

  @Column({ type: 'decimal', precision: 12, scale: 2 })
  amount!: string;

  @Column({ length: 3, default: 'USD' })
  currency!: string;

  @Column({ length: 32, default: 'created' })
  status!: InvoiceStatus;

  @Column({ name: 'workflow_id', type: 'varchar', length: 255, nullable: true })
  workflowId!: string | null;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;
}
