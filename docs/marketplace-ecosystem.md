# Marketplace Ecosystem Guide

## Overview

Marketplace ecosystem extends template lifecycle: discover, submit, review, rate, and apply.

## Endpoints

- `GET /api/v1/marketplace/templates`
- `POST /api/v1/marketplace/templates`
- `POST /api/v1/marketplace/review`
- `POST /api/v1/marketplace/ratings`
- `POST /api/v1/marketplace/apply`

## Review Flow

- `draft` -> `submitted` -> `published`
- `blocked` is supported for risk control.

## Rating and Moderation

- Ratings: 1~5
- Comments can be flagged for moderation.

## Apply

Template apply validates required dependencies before returning success.
